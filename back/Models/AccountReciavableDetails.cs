using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace back.Models;

public class AccountReceivableDetails
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public int AccountReceivableID { get; set; }
    public AccountReceivable AccountReceivable { get; set; }
    public double Amount { get; set; }
}

