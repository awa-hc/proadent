using System.Security.Claims;
using back.Data;
using back.Models;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("[controller]")]
public class UserController : ControllerBase
{
    private readonly ILogger<UserController> _logger;
    private readonly AppDbContext _context;

    private readonly IHttpClientFactory _httpclientFactory;
    public UserController(ILogger<UserController> logger, AppDbContext context, IHttpClientFactory httpclientFactory)
    {
        _logger = logger;
        _context = context;
        _httpclientFactory = httpclientFactory;
    }

    [HttpGet]
    [Authorize(Roles = "admin")]

    public async Task<ActionResult<IEnumerable<User>>> GetUsers()
    {
        return await _context.User.Include(r => r.Role).ToListAsync();
    }

    [HttpGet("{id}")]
    [Authorize(Roles = "admin,user")]
    public async Task<ActionResult<User>> GetUser(int id)
    {
        var user = await _context.User.FindAsync(id);

        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }

        return Ok(user);

    }


    [HttpPost]
    public async Task<ActionResult<User>> RegisterUser([FromBody] RegisterModel request)
    {

        if (string.IsNullOrEmpty(request.Email) || string.IsNullOrEmpty(request.Password) || string.IsNullOrEmpty(request.FullName))
        {
            return BadRequest(new { error = "Email, Password and FullName are required" });
        }
        if (request.Password.Length < 8)
        {
            return BadRequest(new { error = "Password must be at least 8 characters" });
        }
        if (request.Password != request.Password_Confirmation)
        {
            return BadRequest(new { error = "Password and Confirmation must match" });
        }



        var emailExists = await _context.User.FirstOrDefaultAsync(u => u.Email == request.Email);
        if (emailExists != null)
        {
            return BadRequest(new { error = "Email already exists" });
        }

        var ciExists = await _context.User.FirstOrDefaultAsync(u => u.Ci == request.Ci);
        if (ciExists != null)
        {
            return BadRequest(new { error = "CI already exists" });
        }

        if (request.Phone.ToString().Length < 8)
        {
            return BadRequest(new { error = "Phone must be at least 8 characters" });
        }


        request.Password = BCrypt.Net.BCrypt.HashPassword(request.Password);


        User user = new()
        {
            Email = request.Email,
            FullName = request.FullName,
            Password = request.Password,
            Phone = request.Phone,
            Ci = request.Ci,
            BirthDay = request.BirthDay,
            CreatedAt = DateTime.Now.ToUniversalTime(),
            UpdatedAt = DateTime.Now.ToUniversalTime(),
            RoleID = 1
        };

        _context.User.Add(user);
        await _context.SaveChangesAsync();
        if (SendEmailConfirmation(user) == null)
        {
            return BadRequest(new { error = "Failed to send email verification try again later" });
        }
        return CreatedAtAction("GetUser", new { id = user.ID }, user);
    }
    [HttpPut("{id}")]
    [Authorize(Roles = "admin,user")]
    public async Task<ActionResult> PutUser(int id, [FromBody] User request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var user = await _context.User.FindAsync(id);
        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }

        if (string.IsNullOrEmpty(request.Email) || string.IsNullOrEmpty(request.Password) || string.IsNullOrEmpty(request.FullName))
        {
            return BadRequest(new { error = "Email, Password and FullName are required" });
        }
        if (request.Password.Length < 8)
        {
            return BadRequest(new { error = "Password must be at least 8 characters" });
        }

        if (request.Phone.ToString().Length < 8)
        {
            return BadRequest("Phone must be at least 8 characters");
        }
        request.Password = BCrypt.Net.BCrypt.HashPassword(request.Password);
        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        user = request;

        _context.Entry(user).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(user);
    }


    [HttpDelete("{id}")]
    [Authorize(Roles = "admin")]
    public async Task<ActionResult<User>> DeleteUser(int id)
    {
        var user = await _context.User.FindAsync(id);
        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }

        _context.User.Remove(user);
        await _context.SaveChangesAsync();
        return Ok(new { message = "User deleted" });
    }


    private async Task<ActionResult> SendEmailConfirmation(User user)
    {
        var email = user.Email;
        var token = BCrypt.Net.BCrypt.HashPassword(user.Email + user.Password);

        var client = _httpclientFactory.CreateClient();
        var url = "https://dc34sk6l-8080.brs.devtunnels.ms/email";
        var payload = new { email = email, token = token };

        try
        {
            var response = await client.PostAsJsonAsync(url, payload);

            if (response.IsSuccessStatusCode)
            {
                return Ok(new { message = "Confirmation email sent" });
            }
            else
            {
                return StatusCode((int)response.StatusCode, new { error = "Failed to send email" });
            }
        }
        catch (HttpRequestException e)
        {
            return BadRequest(new { error = e.Message });
        }
    }

    [HttpGet("verify-email")]
    public async Task<ActionResult> VerifyEmail(string email, string token)
    {
        var user = await _context.User.FirstOrDefaultAsync(u => u.Email == email);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var expectedToken = BCrypt.Net.BCrypt.HashPassword(user.Email + user.Password);
        if (token != expectedToken)
        {
            return BadRequest(new { error = "Invalid token" });
        }

        await _context.SaveChangesAsync();

        return Ok(new { message = "Email verified successfully" });
    }

    [HttpGet("me")]
    [Authorize]
    public async Task<ActionResult<User>> Me()
    {
        var claim = User.FindFirst(ClaimTypes.Authentication);
        var UserCi = claim != null ? claim.Value.ToString() : null;

        if (UserCi == null)
        {
            return BadRequest(new { error = "user ci is not valid" });
        }
        var appointments = await _context.Appointment
                 .Where(a => a.UserCI == UserCi)
                 .Select(a => new
                 {
                     Date = a.Date.ToString("dd/MM/yyyy HH:mm"),
                     a.Status,
                     a.Reason,
                     CreatedAt = ConvertToTimeZone(a.CreatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
                     UpdatedAt = ConvertToTimeZone(a.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm "),
                     a.Code,
                 })
                 .ToListAsync();
        var user = await _context.User.Include(u => u.Role).Where(u => u.Ci == UserCi).Select(u => new
        {
            u.Ci,
            u.Email,
            u.FullName,
            u.Phone,
            CreatedAt = ConvertToTimeZone(u.CreatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            UpdateAt = ConvertToTimeZone(u.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            u.BirthDay,
            role = new
            {
                u.Role.Name,
                u.Role.Description
            },
            appointments

        }).FirstOrDefaultAsync();


        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }

        return Ok(user);

    }


    [HttpPut("update-userrole/{ci}")]
    public async Task<ActionResult> UpdateUserRole(string ci, [FromBody] UpdateUserRoleRequest request)
    {
        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == ci);
        if (user == null)
        {
            return NotFound(new { error = "User not found" });
        }
        if (request.RoleID == 0)
        {
            return BadRequest(new { error = "Role ID is required" });
        }
        var role = await _context.Role.FindAsync(request.RoleID);
        if (role == null)
        {
            return NotFound(new { error = "Role not found" });
        }
        user.RoleID = request.RoleID;
        _context.Entry(user).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(new { message = "User Role updated" });
    }

    [HttpGet("all")]
    public async Task<ActionResult> GetAllUsers()
    {
        var users = await _context.User.Include(r => r.Role).ToListAsync();
        if (users == null)
        {
            return NotFound(new { error = "users not founded" });
        }

        var userfilter = users.Select(user => new
        {
            user.Ci,
            user.Email,
            user.FullName,
            CreatedAt = ConvertToTimeZone(user.CreatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            UpdatedAt = ConvertToTimeZone(user.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            role = user.Role.Name,
            user.BirthDay,
            user.Phone,
        });
        return Ok(userfilter);
    }


    private static DateTime ConvertToTimeZone(DateTime dateTime, string timeZoneId)
    {
        var timeZoneInfo = TimeZoneInfo.FindSystemTimeZoneById(timeZoneId);
        return TimeZoneInfo.ConvertTimeFromUtc(dateTime, timeZoneInfo);
    }

}

