using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("/[controller]")]
public class AppointmentController : ControllerBase
{

    private readonly ILogger<AppointmentController> _logger;
    private readonly AppDbContext _context;
    private readonly IHttpClientFactory _httpclientFactory;
    public AppointmentController(IHttpClientFactory httpClientFactory, ILogger<AppointmentController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
        _httpclientFactory = httpClientFactory;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<Appointment>>> GetAppointments()
    {
        return await _context.Appointment.ToListAsync();
    }

    [HttpPost]
    public async Task<ActionResult> CreateAppointment([FromBody] AppointmentRequest request)
    {

        if (request.Date < DateTime.Now.ToUniversalTime())
        {
            return BadRequest(new { error = "Date must be in the future" });
        }

        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        if (request.Reason.Length > 100)
        {
            return BadRequest(new { error = "Reason must be less than 100 characters" });
        }
        Appointment appointment = new()
        {
            UserID = request.UserID,
            User = user,
            Date = request.Date,
            Code = await GenerateAppointmentCode(),
            Type = "General",
            UpdatedAt = DateTime.Now.ToUniversalTime(),
            CreatedAt = DateTime.Now.ToUniversalTime(),
            Status = "Pending",
            Reason = request.Reason,
        };

        await _context.Appointment.AddAsync(appointment);
        await _context.SaveChangesAsync();
        if (await SendEmailAppointmentCreated(appointment) is BadRequestObjectResult response)
        {
            return response;
        }
        return Ok(new { message = "Appointment created" });
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<Appointment>> GetAppointment(int id)
    {
        var appointment = await _context.Appointment.FindAsync(id);

        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }

        return Ok(appointment);

    }

    [HttpPut("edit/{id}")]
    public async Task<ActionResult> PutAppointment(int id, [FromBody] Appointment request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var appointment = await _context.Appointment.FindAsync(id);
        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }

        if (request.Date < DateTime.Now.ToUniversalTime())
        {
            return BadRequest(new { error = "Date must be in the future" });
        }

        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        if (request.Reason.Length > 100)
        {
            return BadRequest(new { error = "Reason must be less than 100 characters" });
        }

        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        appointment = request;
        _context.Entry(appointment).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(appointment);
    }

    [HttpDelete("delete/{id}")]
    public async Task<ActionResult> DeleteAppointment(int id)
    {
        var appointment = await _context.Appointment.FindAsync(id);
        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }
        _context.Appointment.Remove(appointment);
        await _context.SaveChangesAsync();
        return Ok(new { message = "Appointment deleted" });
    }


    [HttpPut("status/{id}")]
    public async Task<ActionResult> ChangeStatus(int id, [FromBody] Appointment request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var appointment = await _context.Appointment.FindAsync(id);
        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }
        if (request.Status != "Pending" && request.Status != "Confirmed" && request.Status != "Cancelled" && request.Status == "Completed")
        {
            return BadRequest(new { error = "Invalid Status" });
        }
        appointment.Status = request.Status;
        appointment.UpdatedAt = DateTime.Now.ToUniversalTime();
        _context.Entry(appointment).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(appointment);
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

    private async Task<ActionResult> SendEmailAppointmentCreated(Appointment appointment)
    {
        var user = await _context.User.FindAsync(appointment.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        var client = _httpclientFactory.CreateClient();
        var code = appointment.Code;
        var fullName = user.FullName;
        var date = appointment.Date;
        var email = user.Email;
        var url = "http://localhost:8080/appointment-created";
        var payload = new
        {
            email,
            fullName,
            date,
            code
        };

        try
        {
            var response = await client.PostAsJsonAsync(url, payload);
            if (response.IsSuccessStatusCode)
            {
                return Ok(new { error = "Created Appointment email sent" });
            }
            else
            {
                return BadRequest(new { error = "Error sending email" });
            }

        }
        catch
        {
            return BadRequest(new { error = "Error sending email" });
        }

    }


}
