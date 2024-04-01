using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;

namespace back.Models;

public class Role
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public string Name { get; set; }
    public string Description { get; set; }

    // Relación Uno a Muchos: Un Role puede estar asociado a muchos Users

    public ICollection<User>? Users { get; set; }
}
