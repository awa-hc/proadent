using System.Net.Http;
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
        return await _context.User.ToListAsync();
    }

    [HttpGet("{id}")]
    [Authorize(Roles = "admin,users")]
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


        await SendEmailConfirmation(user);

        return CreatedAtAction("GetUser", new { id = user.ID }, user);
    }
    [HttpPut("{id}")]
    [Authorize(Roles = "admin,users")]
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
        request.UpdatedAt = DateTime.Now;
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
        var url = "http://localhost:8080/send-email/";
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




}

