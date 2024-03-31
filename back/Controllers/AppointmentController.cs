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
        var allappointments = await _context.Appointment.Include(r => r.User).ToListAsync();
        var filtredresponse = allappointments.Select(appointment => new
        {

            appointment.Code,
            appointment.Status,
            date = appointment.Date.ToString("dd/MM/yyyy HH:mm"),
            createdAt = ConvertToTimeZone(appointment.CreatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            updatedAt = ConvertToTimeZone(appointment.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            appointment.Reason,
            UserCi = appointment.User.Ci,
            UserPhone = appointment.User.Phone,
            UserEmail = appointment.User.Email,
            UserName = appointment.User.FullName,
        }).ToList();
        return Ok(filtredresponse);

    }

    [HttpPost]
    public async Task<ActionResult> CreateAppointment([FromBody] AppointmentRequest request)
    {

        if (request.Date < DateTime.Now.ToUniversalTime())
        {
            return BadRequest(new { error = "Date must be in the future" });
        }

        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == request.UserCI);
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
            UserCI = request.UserCI,
            User = user,
            Date = request.Date,
            Code = await GenerateAppointmentCode(),
            Type = "General",
            UpdatedAt = DateTime.Now.ToUniversalTime(),
            CreatedAt = DateTime.Now.ToUniversalTime(),
            Status = "pending",
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

    [HttpGet("{code}")]
    public async Task<ActionResult<Appointment>> GetAppointment(string code)
    {
        var appointment = await _context.Appointment.Include(u => u.User).FirstOrDefaultAsync(a => a.Code == code);


        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }

        var response = new
        {
            appointment.Code,
            date = appointment.Date.ToString("dd/MM/yyyy HH:mm"),
            createdAt = ConvertToTimeZone(appointment.CreatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            updatedAt = ConvertToTimeZone(appointment.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            appointment.Status,
            appointment.Reason,
            userEmail = appointment.User.Email,
            userName = appointment.User.FullName,
            userPhone = appointment.User.Phone,
            userCI = appointment.User.Ci,
            userBirthDay = appointment.User.BirthDay,


        };

        return Ok(response);

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

        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == request.UserCI);
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


    [HttpPut("status/{code}")]
    public async Task<ActionResult> ChangeStatus(string code, [FromBody] UpdateAppointmentRequest request)
    {
        var appointment = await _context.Appointment.FirstOrDefaultAsync(a => a.Code == code);
        if (appointment == null)
        {
            return NotFound(new { error = "Appointment not found" });
        }

        if (request.status == appointment.Status)
        {
            return BadRequest(new { error = "Error update status" });
        }


        switch (request.status.ToLower())
        {
            case "pending":
            case "confirmed":
            case "cancelled":
            case "completed":
                appointment.Status = request.status;
                break;
            default:
                return BadRequest(new { error = "invalid status" });
        }
        appointment.UpdatedAt = DateTime.Now.ToUniversalTime();
        try
        {

            if (await SendEmailAppointmentStatusChanged(appointment) is BadRequestObjectResult response)
            {
                return response;
            }
            await _context.SaveChangesAsync();

        }
        catch (DbUpdateConcurrencyException)
        {
            if (!AppointmentExists(appointment.ID))
            {
                return NotFound(new { error = "Appointment not found" });
            }
            else
            {
                throw;
            }

        }
        return Ok(new { message = "Successfully" });
    }
    private bool AppointmentExists(int id)
    {
        return _context.Appointment.Any(a => a.ID == id);
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

    private async Task<ActionResult> SendEmailAppointmentStatusChanged(Appointment appointment)
    {
        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == appointment.UserCI);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var client = _httpclientFactory.CreateClient();
        var status = appointment.Status;
        var date = appointment.Date;
        var email = user.Email;
        var code = appointment.Code;
        var fullName = user.FullName;
        var url = "https://proadentservicess.fly.dev/appointment-status";
        var payload = new
        {
            code,
            status,
            date,
            updatedAt = ConvertToTimeZone(appointment.UpdatedAt, "SA Western Standard Time").ToString("dd/MM/yyyy HH:mm"),
            appointment.Reason,
            email,
            fullName
        };

        try
        {
            var response = await client.PostAsJsonAsync(url, payload);
            if (response.IsSuccessStatusCode)
            {
                return Ok(new { message = "Update status email sent" });
            }
            else
            {
                return BadRequest(new { error = "error sending email" });
            }
        }
        catch
        {
            return BadRequest(new { error = "erorr sending email" });
        }

    }

    private async Task<ActionResult> SendEmailAppointmentCreated(Appointment appointment)
    {
        var user = await _context.User.FirstOrDefaultAsync(u => u.Ci == appointment.UserCI);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }
        var client = _httpclientFactory.CreateClient();
        var code = appointment.Code;
        var fullName = user.FullName;
        var date = appointment.Date;
        var email = user.Email;
        var url = "https://proadentservicess.fly.dev/appointment-created";
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
                return Ok(new { message = "Created Appointment email sent" });
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


    private DateTime ConvertToTimeZone(DateTime dateTime, string timeZoneId)
    {
        var timeZoneInfo = TimeZoneInfo.FindSystemTimeZoneById(timeZoneId);
        return TimeZoneInfo.ConvertTimeFromUtc(dateTime, timeZoneInfo);
    }




}
