using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("Price/[controller]")]
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
    public async Task<ActionResult> CreatePrice([FromBody] Price request)
    {
        var user = await _context.User.FindAsync(request.UserID);
        var procedure = await _context.Procedure.FindAsync(request.ProcedureID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
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

        request.CreatedAt = DateTime.Now;
        request.UpdatedAt = DateTime.Now;

        Price price = new()
        {
            UserID = request.UserID,
            User = user,
            ProcedureID = request.ProcedureID,
            Procedure = procedure,
            AppointmentDays = request.AppointmentDays,
            TotalPrice = request.TotalPrice,
            CreatedAt = request.CreatedAt,
            UpdatedAt = request.UpdatedAt
        };

        await _context.Price.AddAsync(price);
        await _context.SaveChangesAsync();
        return Ok(price);
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

        var user = await _context.User.FindAsync(request.UserID);
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
    public async Task<ActionResult> ChangeStatus(int id, [FromBody] Price request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var price = await _context.Price.FindAsync(id);
        if (price == null)
        {
            return NotFound(new { error = "Price not found" });
        }
        if (request.Status != "Pending" && request.Status != "Confirmed" && request.Status != "Cancelled")
        {
            return BadRequest(new { error = "Invalid Status" });
        }
        price.Status = request.Status;
        if (request.Status == "Confirmed")
        {
            string codeAccount = await GenerateAccountReceivableCode();
            AccountReceivable accountReceivable = new AccountReceivable
            {
                UserID = request.UserID,
                Code = codeAccount,
                AppointmentDays = request.AppointmentDays,
                ProceduresDescription = price.Procedure.Description,
                TotalPrice = request.TotalPrice,
                Balance = request.TotalPrice,
                Status = "Pending",
                CreatedAt = DateTime.Now,
                UpdatedAt = DateTime.Now
            };

            await _context.AccountReceivable.AddAsync(accountReceivable);

            int[] paymentPlan = new int[1];
            paymentPlan = CalculatePaymentPlan(request.TotalPrice, request.AppointmentDays);
            var user = await _context.User.FindAsync(request.UserID);
            for (int i = 0; i < request.AppointmentDays; i++)
            {
                AccountReceivableDetail accountReceivableDetail = new()
                {
                    AccountReceivableID = accountReceivable.ID,
                    AccountReceivable = accountReceivable,
                    Amount = paymentPlan[i],
                    Status = "Pending",
                    CreatedAt = DateTime.Now,
                    UpdatedAt = DateTime.Now
                };
                await _context.AccountReceivableDetail.AddAsync(accountReceivableDetail);

                Appointment appointment = new()
                {
                    UserID = request.UserID,
                    User = user,
                    Date = DateTime.Now,
                    Code = await GenerateAppointmentCode(),
                    Type = "General",
                    UpdatedAt = DateTime.Now,
                    CreatedAt = DateTime.Now,
                    Status = "Pending",
                    Reason = price.Procedure.Description,
                };
                await _context.Appointment.AddAsync(appointment);
            }
        }
        price.UpdatedAt = DateTime.Now;
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
