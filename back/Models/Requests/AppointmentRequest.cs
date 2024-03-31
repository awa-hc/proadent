namespace back.Models;

public class AppointmentRequest
{
    public string UserCI { get; set; }
    public string? Type { get; set; }
    public DateTime Date { get; set; }
    public required string Reason { get; set; }
}