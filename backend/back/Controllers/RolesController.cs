using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("[controller]")]
public class RoleController : ControllerBase
{
    private readonly ILogger<RoleController> _logger;
    private readonly AppDbContext _context;
    public RoleController(ILogger<RoleController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }
    [HttpPost]
    public async Task<ActionResult<Role>> RegisterRole([FromBody] Role request)
    {

        var existingRole = await _context.Role.FirstOrDefaultAsync(r => r.Name == request.Name);
        if (existingRole != null)
        {
            return BadRequest(new { error = "Role al ready exists" });
        }
        if (string.IsNullOrEmpty(request.Name))
        {
            return BadRequest(new { error = "Role name are required" });
        }
        if (string.IsNullOrEmpty(request.Description))
        {
            return BadRequest(new { error = "Role Description are required" });
        }


        Role role = new()
        {
            Name = request.Name,
            Description = request.Description
        };

        _context.Role.Add(role);
        await _context.SaveChangesAsync();
        return Ok(new { message = "Role Created successfully" });
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<Role>>> GetRoles()
    {
        return await _context.Role.ToListAsync();

    }

}