namespace back.Models;

public class CreateAccountReceivableRequest
{
    public string UserCI { get; set; }
    public string ProceduresDescription { get; set; }

    public int AppointmentDays { get; set; }
    public decimal TotalPrice { get; set; }
    public decimal Balance { get; set; }


}