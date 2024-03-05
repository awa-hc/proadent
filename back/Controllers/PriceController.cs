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

}
