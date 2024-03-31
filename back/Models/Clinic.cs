namespace back.Models;

using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class Clinic
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public string UserCI { get; set; }
    public User User { get; set; }
    public string AppointmentCode { get; set; }
    public Appointment Appointment { get; set; }
    public int ProcedureID { get; set; }
    public Procedure Procedure { get; set; }
    public int PriceID { get; set; }
    public Price Price { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
}
