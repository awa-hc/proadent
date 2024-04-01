namespace back.Models;

using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text.Json.Serialization;

public class User
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
    public string Email { get; set; }
    public required string FullName { get; set; }
    public string Password { get; set; }
    public string Phone { get; set; }
    public string Ci { get; set; }
    public DateOnly BirthDay { get; set; }
    public int RoleID { get; set; }
    public Role Role { get; set; }
}
