using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using back.Data;
using back.Models;
using back.Models.Auth;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;

namespace back.Controllers;

[ApiController]
[Route("/[controller]")]
public class AuthController : ControllerBase
{
    private readonly AppDbContext _context;
    private readonly IConfiguration _configuration;
    public AuthController(AppDbContext context, IConfiguration configuration)
    {
        _context = context;
        _configuration = configuration;
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginModel request)
    {

        if (request.Email == null && request.Ci == null)
        {
            return BadRequest(new { error = "Email or CI is required" });
        }
        if (request.Password == null)
        {
            return BadRequest(new { error = "Password is required" });
        }

        var user = await _context.User.Include(r => r.Role).FirstOrDefaultAsync(u => u.Email == request.Email || u.Ci == request.Ci);
        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }
        if (!BCrypt.Net.BCrypt.Verify(request.Password, user.Password))
        {
            return BadRequest(new { error = "Invalid password" });
        }
        if (user.Role == null)
        {
            return BadRequest(new { error = "User has no role" });
        }
        if (user.Role.Name == "disabled")
        {
            return BadRequest(new { error = "User is disabled" });
        }
        if (user.FullName == null)
        {
            return BadRequest(new { error = "User has no name" });
        }

        var claims = new[]{
            new Claim(ClaimTypes.Name, user.FullName ),
            new Claim(ClaimTypes.Role, user.Role.Name ),
            new Claim(ClaimTypes.NameIdentifier, user.ID.ToString() ),
            new Claim(ClaimTypes.Authentication, user.Ci)
        };

        var keyString = _configuration["Jwt:Key"]; // Obtiene el valor de la configuración
        var key = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(keyString)); // Convierte la clave a bytes
        var issuer = _configuration["Jwt:Issuer"];
        var audience = _configuration["Jwt:Audience"];
        var creds = new SigningCredentials(key, SecurityAlgorithms.HmacSha256);

        var token = new JwtSecurityToken(
            issuer: issuer,
            audience: audience,
            claims: claims,
            expires: DateTime.Now.AddDays(30),
            signingCredentials: creds
        );



        return Ok(new { token = new JwtSecurityTokenHandler().WriteToken(token) });

    }

}
