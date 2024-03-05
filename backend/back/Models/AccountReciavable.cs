using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace back.Models;
public class AccountReceivable
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public int UserID { get; set; }
    public User User { get; set; }
    public int AppointmentDays { get; set; }
    public string ProceduresDescription { get; set; }
    public decimal TotalPrice { get; set; }
    public decimal Balance { get; set; }

    // Relación Uno a Muchos: Un AccountReceivable puede tener muchos Details
    public ICollection<AccountReceivableDetails> Details { get; set; }
}
