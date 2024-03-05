namespace back.Models;

using System;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class Appointment
{
    [Key]
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public int ID { get; set; }
    public int UserID { get; set; }
    public User User { get; set; }
    public string Code { get; set; }
    public string Type { get; set; }
    public DateTime Date { get; set; }
    public string Reason { get; set; }
    public string Status { get; set; }
    public DateTime CreatedAt { get; set; }
    public DateTime UpdatedAt { get; set; }
    // public string Observations { get; set; }
    // public string Doctor { get; set; }
    // public string Specialty { get; set; }
    // public string PaymentMethod { get; set; }
    // public string PaymentStatus { get; set; }
    // public string PaymentReference { get; set; }
    // public string PaymentAmount { get; set; }
    // public string PaymentCurrency { get; set; }
    // public string PaymentDescription { get; set; }
    // public string PaymentObservations { get; set; }
    // public string PaymentProvider { get; set; }
    // public string PaymentProviderID { get; set; }
    // public string PaymentProviderStatus { get; set; }
    // public string PaymentProviderMessage { get; set; }
    // public string PaymentProviderReference { get; set; }
    // public string PaymentProviderAmount { get; set; }
    // public string PaymentProviderCurrency { get; set; }
    // public string PaymentProviderDescription { get; set; }
    // public string PaymentProviderObservations { get; set; }
    // public string PaymentProviderCreated { get; set; }
    // public string PaymentProviderUpdated { get; set; }
    // public string PaymentProviderDeleted { get; set; }
    // public string PaymentProviderStatusReason { get; set; }
    // // Relación Uno a Uno: Un Appointment tiene un AccountReceivable
    // public AccountReceivable AccountReceivable { get; set; }
    // public int AccountReceivableID { get; set; }
    // public string AccountReceivableStatus { get; set; }

    // // Relación Uno a Uno: Un Appointment tiene un Price
    // public Price Price { get; set; }

}
