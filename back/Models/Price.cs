using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace back.Models
{
    public class Price
    {
        [Key]
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public int ID { get; set; }
        public int UserID { get; set; }
        public User User { get; set; }
        public int ProcedureID { get; set; }
        public Procedure Procedure { get; set; }
        public int AppointmentDays { get; set; }
        public decimal TotalPrice { get; set; }
    }

}