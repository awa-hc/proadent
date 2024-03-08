using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("AccountReceivable/[controller]")]


public class AccountReceivableController : ControllerBase
{
    private readonly ILogger<AccountReceivableController> _logger;
    private readonly AppDbContext _context;
    public AccountReceivableController(ILogger<AccountReceivableController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<AccountReceivable>>> GetAccountReceivables()
    {
        return await _context.AccountReceivable.ToListAsync();
    }

    [HttpPost]
    public async Task<ActionResult> CreateAccountReceivable([FromBody] AccountReceivable request)
    {

        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        if (request.ProceduresDescription == null)
        {
            return BadRequest(new { error = "ProceduresDescription is required" });
        }
        if (request.AppointmentDays <= 0)
        {
            return BadRequest(new { error = "AppointmentDays must be greater than 0" });
        }
        if (request.TotalPrice <= 0)
        {
            return BadRequest(new { error = "TotalPrice must be greater than 0" });
        }

        request.Balance = request.TotalPrice;
        request.Status = "Pending";
        request.CreatedAt = DateTime.Now.ToUniversalTime();
        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        request.Code = await GenerateAccountReceivableCode();


        AccountReceivable accountReceivable = new()
        {
            UserID = request.UserID,
            User = user,
            Code = request.Code,
            AppointmentDays = request.AppointmentDays,
            ProceduresDescription = request.ProceduresDescription,
            TotalPrice = request.TotalPrice,
            Balance = request.Balance,
            Status = request.Status,
            CreatedAt = request.CreatedAt,
            UpdatedAt = request.UpdatedAt
        };

        await _context.AccountReceivable.AddAsync(accountReceivable);
        await _context.SaveChangesAsync();
        return Ok(accountReceivable);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<AccountReceivable>> GetAccountReceivable(int id)
    {
        var AccountReceivable = await _context.AccountReceivable.FindAsync(id);

        if (AccountReceivable == null)
        {
            return NotFound(new { error = "Account Receivable not found" });
        }

        return Ok(AccountReceivable);

    }

    [HttpPut("edit/{id}")]
    public async Task<ActionResult> PutAccountReceivable(int id, [FromBody] AccountReceivable request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var accountReceivable = await _context.AccountReceivable.FindAsync(id);
        if (accountReceivable == null)
        {
            return BadRequest(new { error = "AccountReceivable not found" });
        }
        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        if (request.ProceduresDescription == null)
        {
            return BadRequest(new { error = "ProceduresDescription is required" });
        }
        if (request.AppointmentDays <= 0)
        {
            return BadRequest(new { error = "AppointmentDays must be greater than 0" });
        }
        if (request.TotalPrice <= 0)
        {
            return BadRequest(new { error = "TotalPrice must be greater than 0" });
        }

        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        accountReceivable = request;
        _context.Entry(accountReceivable).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(accountReceivable);
    }

    [HttpDelete("delete/{id}")]
    public async Task<ActionResult> DeleteAccountReceivable(int id)
    {
        var accountReceivable = await _context.AccountReceivable.FindAsync(id);
        if (accountReceivable == null)
        {
            return NotFound(new { error = "AccountReceivable not found" });
        }
        _context.AccountReceivable.Remove(accountReceivable);
        await _context.SaveChangesAsync();
        return Ok(new { message = "AccountReceivable deleted" });
    }


    [HttpPut("status/{id}")]
    public async Task<ActionResult> ChangeStatus(int id, [FromBody] AccountReceivable request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var accountReceivable = await _context.AccountReceivable.FindAsync(id);
        if (accountReceivable == null)
        {
            return NotFound(new { error = "AccountReceivable not found" });
        }
        if (request.Status != "Pending" && request.Status != "Completed" && request.Status != "Cancelled" && request.Status == "Completed")
        {
            return BadRequest(new { error = "Invalid Status" });
        }
        accountReceivable.Status = request.Status;
        accountReceivable.UpdatedAt = DateTime.Now.ToUniversalTime();
        _context.Entry(accountReceivable).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(accountReceivable);
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
}
