namespace back.Models;

using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class Clinic
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public int UserID { get; set; }
    public User User { get; set; }
    public int AppointmentID { get; set; }
    public Appointment Appointment { get; set; }
    public int ProcedureID { get; set; }
    public Procedure Procedure { get; set; }
    public int PriceID { get; set; }
    public Price Price { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
}
