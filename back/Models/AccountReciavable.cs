namespace back.Models;

using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class AccountReceivable
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public int UserID { get; set; }
    public User User { get; set; }
    public string Code { get; set; }
    public int AppointmentDays { get; set; }
    public string ProceduresDescription { get; set; }
    public decimal TotalPrice { get; set; }
    public decimal Balance { get; set; }
    public string Status { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }

    // Relación Uno a Muchos: Un AccountReceivable puede tener muchos Details
    public ICollection<AccountReceivableDetail> Details { get; set; }
}
