using back.Data;
using back.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace back.Controllers;

[ApiController]
[Route("Clinic/[controller]")]
public class ClinicController : ControllerBase
{

    private readonly ILogger<ClinicController> _logger;
    private readonly AppDbContext _context;
    public ClinicController(ILogger<ClinicController> logger, AppDbContext context)
    {
        _logger = logger;
        _context = context;
    }

    [HttpGet]
    public async Task<ActionResult<IEnumerable<Clinic>>> GetClinics()
    {
        return await _context.Clinic.ToListAsync();
    }

    [HttpPost]
    public async Task<ActionResult> CreateClinic([FromBody] Clinic request)
    {

        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var appointment = await _context.Appointment.FindAsync(request.AppointmentID);
        if (appointment == null)
        {
            return BadRequest(new { error = "Appointment not found" });
        }

        var procedure = await _context.Procedure.FindAsync(request.ProcedureID);
        if (procedure == null)
        {
            return BadRequest(new { error = "Procedure not found" });
        }

        var price = await _context.Price.FindAsync(request.PriceID);
        if (price == null)
        {
            return BadRequest(new { error = "Price not found" });
        }

        request.CreatedAt = DateTime.Now.ToUniversalTime();
        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        //request.Code = await GenerateClinicCode(request);


        Clinic clinic = new()
        {
            UserID = request.UserID,
            User = user,
            AppointmentID = request.AppointmentID,
            Appointment = appointment,
            ProcedureID = request.ProcedureID,
            Procedure = procedure,
            PriceID = request.PriceID,
            Price = price,
            //Code = request.Code,
            CreatedAt = request.CreatedAt,
            UpdatedAt = request.UpdatedAt
        };

        await _context.Clinic.AddAsync(clinic);
        await _context.SaveChangesAsync();
        return Ok(clinic);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<Clinic>> GetClinic(int id)
    {
        var clinic = await _context.Clinic.FindAsync(id);

        if (clinic == null)
        {
            return NotFound(new { error = "Clinic not found" });
        }

        return Ok(clinic);

    }

    [HttpPut("edit/{id}")]
    public async Task<ActionResult> PutClinic(int id, [FromBody] Clinic request)
    {
        if (id != request.ID)
        {
            return BadRequest(new { error = "Invalid ID" });
        }
        var clinic = await _context.Clinic.FindAsync(id);
        if (clinic == null)
        {
            return NotFound(new { error = "Clinic not found" });
        }

        var user = await _context.User.FindAsync(request.UserID);
        if (user == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var appointment = await _context.Appointment.FindAsync(request.AppointmentID);
        if (appointment == null)
        {
            return BadRequest(new { error = "Appointment not found" });
        }

        var procedure = await _context.Procedure.FindAsync(request.ProcedureID);
        if (procedure == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        var price = await _context.Price.FindAsync(request.PriceID);
        if (price == null)
        {
            return BadRequest(new { error = "User not found" });
        }

        request.UpdatedAt = DateTime.Now.ToUniversalTime();
        clinic = request;
        _context.Entry(clinic).State = EntityState.Modified;
        await _context.SaveChangesAsync();
        return Ok(clinic);
    }

    [HttpDelete("delete/{id}")]
    public async Task<ActionResult> DeleteClinic(int id)
    {
        var clinic = await _context.Clinic.FindAsync(id);
        if (clinic == null)
        {
            return NotFound(new { error = "Clinic not found" });
        }
        _context.Clinic.Remove(clinic);
        await _context.SaveChangesAsync();
        return Ok(new { message = "Clinic deleted" });
    }

    // private async Task<string> GenerateClinicCode(Clinic request)
    // {
    //     var lastClinic = await _context.Clinic.OrderByDescending(a => a.ID).FirstOrDefaultAsync();
    //     if (lastClinic == null)
    //     {
    //         return "C-0001";
    //     }
    //     string[] code = lastClinic.code.Split("-");
    //     int number = int.Parse(code[1]);
    //     number++;
    //     return "C-" + number.ToString("D4");
    // }
}
