
using back.Models;
using Microsoft.EntityFrameworkCore;
using System;
namespace back.Data
{
    public class AppDbContext : DbContext
    {
        public AppDbContext(DbContextOptions<AppDbContext> options)
            : base(options)
        {
        }

        public DbSet<User> User { get; set; }
        public DbSet<Role> Role { get; set; }
        public DbSet<Appointment> Appointment { get; set; }
        public DbSet<Procedure> Procedure { get; set; }
        public DbSet<Price> Price { get; set; }
        public DbSet<Clinic> Clinic { get; set; }
        public DbSet<AccountReceivable> AccountReceivable { get; set; }
        public DbSet<AccountReceivableDetail> AccountReceivableDetail { get; set; }

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {

        }
    }
}
