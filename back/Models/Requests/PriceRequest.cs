namespace back.Models;


public class PriceRequest
{
    public string UserCI { get; set; }

    public int ProcedureID { get; set; }

    public int AppointmentDays { get; set; }
    public decimal TotalPrice { get; set; }

    public string Status { get; set; }

}