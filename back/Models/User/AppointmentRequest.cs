namespace back.Models;

public class AppointmentRequest
{
    public int UserID { get; set; }
    public string? Type { get; set; }
    public DateTime Date { get; set; }
    public required string Reason { get; set; }
}