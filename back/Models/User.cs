namespace back.Models;

using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

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

    // Relación Uno a Muchos: Un usuario puede tener muchos Appointments
    public ICollection<Appointment> Appointments { get; set; }

    // Relación Uno a Muchos: Un usuario puede tener muchas Clinics
    public ICollection<Clinic> Clinics { get; set; }

    // Relación Uno a Muchos: Un usuario puede tener muchos Prices
    public ICollection<Price> Prices { get; set; }

    // Relación Uno a Muchos: Un usuario puede tener muchos AccountReceivables
    public ICollection<AccountReceivable> AccountReceivables { get; set; }

    // Relación Uno a Uno: Un usuario tiene un Role
    public int RoleID { get; set; }
    public Role Role { get; set; }
}
