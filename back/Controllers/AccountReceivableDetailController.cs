using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("AccountReceivableDetail/[controller]")]


public class AccountReceivableDetailController : ControllerBase
{
    private readonly ILogger<AccountReceivableDetailController> _logger;
    private readonly AppDbContext _context;
    public AccountReceivableDetailController(ILogger<AccountReceivableDetailController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<AccountReceivableDetail>>> GetAccountReceivableDetail()
    {
        return await _context.AccountReceivableDetail.ToListAsync();
    }

    [HttpPost]
    public async Task<ActionResult> CreateAccountReceivableDetail([FromBody] AccountReceivableDetail request)
    {

        var accountReceivable = await _context.AccountReceivable.FindAsync(request.AccountReceivableID);
        if (accountReceivable == null)
        {
            return BadRequest(new { error = "AccountReceivable not found" });
        }
        if (request.Amount <= 0)
        {
            return BadRequest(new { error = "Amount must be greater than 0" });
        }

        request.Status = "Pending";
        request.CreatedAt = DateTime.Now.ToUniversalTime();
        request.UpdatedAt = DateTime.Now.ToUniversalTime();


        AccountReceivableDetail AccountReceivableDetail = new()
        {
            AccountReceivableID = request.AccountReceivableID,
            AccountReceivable = accountReceivable,
            Amount = request.Amount,
            Status = request.Status,
            CreatedAt = request.CreatedAt,
            UpdatedAt = request.UpdatedAt
        };

        await _context.AccountReceivableDetail.AddAsync(AccountReceivableDetail);
        await _context.SaveChangesAsync();
        return Ok(AccountReceivableDetail);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<AccountReceivableDetail>> GetAccountReceivableDetail(int id)
    {
        var AccountReceivableDetail = await _context.AccountReceivableDetail.FindAsync(id);

        if (AccountReceivableDetail == null)
        {
            return NotFound(new { error = "AccountReceivableDetail not found" });
        }

        return Ok(AccountReceivableDetail);

    }

    [HttpPut("edit/{id}")]
    public async Task<ActionResult> PutAccountReceivableDetail(int id, [FromBody] AccountReceivableDetail request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var accountReceivableDetail = await _context.AccountReceivableDetail.FindAsync(id);
        if (accountReceivableDetail == null)
        {
            return BadRequest(new { error = "AccountReceivableDetail not found" });
        }
        var accountReceivable = await _context.AccountReceivable.FindAsync(request.AccountReceivableID);
        if (accountReceivable == null)
        {
            return BadRequest(new { error = "AccountReceivable not found" });
        }
        if (request.Amount <= 0)
        {
            return BadRequest(new { error = "TotalPrice must be greater than 0" });
        }

        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        accountReceivableDetail = request;
        _context.Entry(accountReceivableDetail).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(accountReceivableDetail);
    }

    [HttpDelete("delete/{id}")]
    public async Task<ActionResult> DeleteAccountReceivableDetail(int id)
    {
        var AccountReceivableDetail = await _context.AccountReceivableDetail.FindAsync(id);
        if (AccountReceivableDetail == null)
        {
            return NotFound(new { error = "AccountReceivableDetail not found" });
        }
        _context.AccountReceivableDetail.Remove(AccountReceivableDetail);
        await _context.SaveChangesAsync();
        return Ok(new { message = "AccountReceivableDetail deleted" });
    }


    [HttpPut("status/{id}")]
    public async Task<ActionResult> ChangeStatus(int id, [FromBody] AccountReceivableDetail request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var accountReceivableDetail = await _context.AccountReceivableDetail.FindAsync(id);
        if (accountReceivableDetail == null)
        {
            return NotFound(new { error = "AccountReceivableDetail not found" });
        }
        if (request.Status != "Pending" && request.Status != "Completed" && request.Status != "Cancelled" && request.Status == "Completed")
        {
            return BadRequest(new { error = "Invalid Status" });
        }
        accountReceivableDetail.Status = request.Status;
        accountReceivableDetail.UpdatedAt = DateTime.Now.ToUniversalTime();
        _context.Entry(accountReceivableDetail).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(accountReceivableDetail);
    }

}
