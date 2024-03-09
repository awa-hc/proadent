namespace back.Models;

using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

public class UpdateUserRoleRequest
{
    public int RoleID { get; set; }
}