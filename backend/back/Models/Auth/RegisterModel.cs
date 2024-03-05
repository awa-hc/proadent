namespace back.Models;

public class RegisterModel
{
    public string Email { get; set; }
    public string FullName { get; set; }
    public string Password { get; set; }
    public string Password_Confirmation { get; set; }
    public string Phone { get; set; }
    public string Ci { get; set; }
    public DateOnly BirthDay { get; set; }
}