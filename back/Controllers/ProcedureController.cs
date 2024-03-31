using System.Security.Cryptography.X509Certificates;
using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;


[ApiController]
[Route("[controller]")]
public class ProcedureController : ControllerBase
{
    private readonly ILogger<ProcedureController> _logger;
    private readonly AppDbContext _context;

    public ProcedureController(ILogger<ProcedureController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }

    [HttpPost]
    public async Task<ActionResult<Procedure>> RegisterProcedure(Procedure request)
    {
        if (string.IsNullOrWhiteSpace(request.Name) || string.IsNullOrEmpty(request.Description))
        {
            return BadRequest(new { error = "Name and description are required" });
        }

        if (request.SuggestedPrice < 0)
        {
            return BadRequest(new { error = "SuggestedPrice must be greater than 0" });
        }

        var existingprocedure = await _context.Procedure.FirstOrDefaultAsync(u => u.Name == request.Name);
        if (existingprocedure != null)
        {
            return BadRequest(new { error = "Procedure already exists" });
        }

        _context.Procedure.Add(request);
        await _context.SaveChangesAsync();
        return Ok(new { message = "Procedure created successfully" });
    }
    [HttpGet]
    public async Task<ActionResult<IEnumerable<Procedure>>> GetProcedures()
    {
        return await _context.Procedure.ToListAsync();
    }


}