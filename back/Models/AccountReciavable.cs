namespace back.Models;

using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class AccountReceivable
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public string UserCI { get; set; }
    public User User { get; set; }
    public string Code { get; set; }
    public int AppointmentDays { get; set; }
    public string ProceduresDescription { get; set; }
    public decimal TotalPrice { get; set; }
    public decimal Balance { get; set; }
    public string Status { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }

    public ICollection<AccountReceivableDetail> Details { get; set; }
}
