using System.Data;
using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("[controller]")]
public class PriceController : ControllerBase
{
    private readonly ILogger<PriceController> _logger;
    private readonly AppDbContext _context;
    public PriceController(ILogger<PriceController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<Price>>> GetPrices()
    {
        return await _context.Price.ToListAsync();
    }

    [HttpPost]
    public async Task<ActionResult> CreatePrice([FromBody] PriceRequest request)
    {
        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == request.UserCI);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        var procedure = await _context.Procedure.FindAsync(request.ProcedureID);
        if (procedure == null)
        {
            return BadRequest(new { error = "Procedure not found" });
        }
        if (request.AppointmentDays <= 0)
        {
            return BadRequest(new { error = "AppointmentDays must be greater than 0" });
        }
        if (request.TotalPrice <= 0)
        {
            return BadRequest(new { error = "TotalPrice must be greater than 0" });
        }
        Price price = new()
        {
            UserCI = request.UserCI,
            User = user,
            ProcedureID = request.ProcedureID,
            Procedure = procedure,
            Status = request.Status,
            AppointmentDays = request.AppointmentDays,
            TotalPrice = request.TotalPrice,
            CreatedAt = DateTime.Now.ToUniversalTime(),
            UpdatedAt = DateTime.Now.ToUniversalTime()
        };

        await _context.Price.AddAsync(price);
        await _context.SaveChangesAsync();
        return Ok(new { message = "price created successfully id:", price.ID });
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<Price>> GetPrice(int id)
    {
        var price = await _context.Price.FindAsync(id);

        if (price == null)
        {
            return NotFound(new { error = "Price not found" });
        }

        return Ok(price);

    }
    [HttpGet("user/{ci}")]
    public async Task<ActionResult<Price>> GetUserPrice(string ci)
    {
        // Validar que el parámetro ci se proporciona
        if (string.IsNullOrWhiteSpace(ci))
        {
            return BadRequest(new { error = "Cédula de identidad no proporcionada." });
        }

        // Buscar el usuario por cédula
        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == ci);
        if (user == null)
        {
            return BadRequest(new { error = "Usuario no encontrado." });
        }

        var userPrices = await _context.Price.Where(p => p.UserCI == user.Ci).Include(u => u.User).Include(p => p.Procedure).ToListAsync();

        var userresponse = userPrices.Select(Price => new
        {
            Price.ID,
            userCi = Price.UserCI,
            userEmail = Price.User.Email,
            userName = Price.User.FullName,
            procedure = Price.Procedure.Name,
            procedureDetails = Price.Procedure.Description,
            Price.CreatedAt,
            Price.UpdatedAt,
            Price.AppointmentDays,
            Price.Status
        });
        return Ok(userresponse);
    }



    [HttpPut("edit/{id}")]
    public async Task<ActionResult> PutPrice(int id, [FromBody] Price request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var price = await _context.Price.FindAsync(id);
        if (price == null)
        {
            return BadRequest(new { error = "Price not found" });
        }

        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == request.UserCI);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var procedure = await _context.Procedure.FindAsync(request.ProcedureID);
        if (procedure == null)
        {
            return BadRequest(new { error = "Procedure not found" });
        }

        if (request.AppointmentDays <= 0)
        {
            return BadRequest(new { error = "AppointmentDays must be greater than 0" });
        }
        if (request.TotalPrice <= 0)
        {
            return BadRequest(new { error = "TotalPrice must be greater than 0" });
        }

        request.UpdatedAt = DateTime.Now;
        price = request;
        _context.Entry(price).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(price);
    }

    [HttpDelete("delete/{id}")]
    public async Task<ActionResult> DeletePrice(int id)
    {
        var price = await _context.Price.FindAsync(id);
        if (price == null)
        {
            return NotFound(new { error = "Price not found" });
        }
        _context.Price.Remove(price);
        await _context.SaveChangesAsync();
        return Ok(new { message = "Price deleted" });
    }

    [HttpPut("change-status/{id}")]
    public async Task<ActionResult> ChangeStatus(int id, [FromBody] UpdatePriceStatus request)
    {
        var price = await _context.Price.Include(p => p.Procedure).FirstOrDefaultAsync(p => p.ID == id);

        if (price == null)
        {
            return NotFound(new { error = "Price not found" });
        }
        var priceuser = await _context.User.FirstOrDefaultAsync(u => u.Ci == price.UserCI);
        if (priceuser == null)
        {
            return BadRequest(new { error = "user not found" });
        }

        if (request.Status != "pending" && request.Status != "confirmed" && request.Status != "cancelled")
        {
            return BadRequest(new { error = "Invalid Status" });
        }

        price.Status = request.Status;
        if (request.Status == "confirmed")
        {
            string codeAccount = await GenerateAccountReceivableCode();
            AccountReceivable accountReceivable = new()
            {
                UserCI = price.UserCI,
                Code = codeAccount,
                User = priceuser,
                AppointmentDays = price.AppointmentDays,
                ProceduresDescription = price.Procedure.Description,
                TotalPrice = price.TotalPrice,
                Balance = price.TotalPrice,
                Status = "pending",
                CreatedAt = DateTime.Now.ToUniversalTime(),
                UpdatedAt = DateTime.Now.ToUniversalTime()
            };

            await _context.AccountReceivable.AddAsync(accountReceivable);

            int[] paymentPlan = CalculatePaymentPlan(price.TotalPrice, price.AppointmentDays);
            var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == price.UserCI);
            for (int i = 0; i < price.AppointmentDays; i++)
            {
                AccountReceivableDetail accountReceivableDetail = new()
                {
                    AccountReceivableCode = accountReceivable.Code,
                    AccountReceivable = accountReceivable,
                    Amount = paymentPlan[i],
                    Status = "pending",
                    CreatedAt = DateTime.Now.ToUniversalTime(),
                    UpdatedAt = DateTime.Now.ToUniversalTime()
                };
                await _context.AccountReceivableDetail.AddAsync(accountReceivableDetail);

                Appointment appointment = new()
                {
                    UserCI = price.UserCI,
                    Date = DateTime.Now.ToUniversalTime().AddDays(i + 1),
                    Code = await GenerateAppointmentCode(),
                    Type = "General",
                    UpdatedAt = DateTime.Now.ToUniversalTime(),
                    CreatedAt = DateTime.Now.ToUniversalTime(),
                    Status = "pending",
                    Reason = price.Procedure.Description,
                };
                await _context.Appointment.AddAsync(appointment);
            }
        }
        price.UpdatedAt = DateTime.Now.ToUniversalTime();
        _context.Entry(price).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(price);
    }

    private async Task<string> GenerateAccountReceivableCode()
    {
        var lastAccountReceivable = await _context.AccountReceivable.OrderByDescending(a => a.ID).FirstOrDefaultAsync();
        if (lastAccountReceivable == null)
        {
            return "AR-0001";
        }
        string[] code = lastAccountReceivable.Code.Split("-");
        int number = int.Parse(code[1]);
        number++;
        return "AR-" + number.ToString("D4");
    }

    private async Task<string> GenerateAppointmentCode()
    {
        var lastAppointment = await _context.Appointment.OrderByDescending(a => a.ID).FirstOrDefaultAsync();
        if (lastAppointment == null)
        {
            return "AP-0001";
        }
        string[] code = lastAppointment.Code.Split("-");
        int number = int.Parse(code[1]);
        number++;
        return "AP-" + number.ToString("D4");
    }

    private int[] CalculatePaymentPlan(decimal totalAmount, int numberOfMonths)
    {
        decimal monthlyAmount = Math.Ceiling(totalAmount / numberOfMonths);

        int[] paymentPlan = new int[numberOfMonths];

        decimal remainingAmount = totalAmount;
        for (int i = 0; i < numberOfMonths; i++)
        {
            if (i == numberOfMonths - 1)
            {
                paymentPlan[i] = (int)remainingAmount;
            }
            else
            {
                paymentPlan[i] = (int)Math.Min(monthlyAmount, remainingAmount);
                remainingAmount -= paymentPlan[i];
            }
        }
        return paymentPlan;
    }


}
