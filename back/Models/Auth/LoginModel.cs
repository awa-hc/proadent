namespace back.Models.Auth
{
    public class LoginModel
    {
        public string? Email { get; set; }
        public string? Ci { get; set; }
        public required string Password { get; set; }
    }
}