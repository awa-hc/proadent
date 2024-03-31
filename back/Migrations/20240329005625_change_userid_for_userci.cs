using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace back.Migrations
{
    /// <inheritdoc />
    public partial class change_userid_for_userci : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<int>(
                name: "SuggestedPrice",
                table: "Procedure",
                type: "integer",
                nullable: false,
                defaultValue: 0);

            migrationBuilder.AddColumn<string>(
                name: "Status",
                table: "Price",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "UserCI",
                table: "Price",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<DateTime>(
                name: "CreatedAt",
                table: "Clinic",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.AddColumn<DateTime>(
                name: "UpdatedAt",
                table: "Clinic",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.AddColumn<string>(
                name: "UserCI",
                table: "Appointment",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<string>(
                name: "UserCI",
                table: "AccountReceivable",
                type: "text",
                nullable: false,
                defaultValue: "");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "SuggestedPrice",
                table: "Procedure");

            migrationBuilder.DropColumn(
                name: "Status",
                table: "Price");

            migrationBuilder.DropColumn(
                name: "UserCI",
                table: "Price");

            migrationBuilder.DropColumn(
                name: "CreatedAt",
                table: "Clinic");

            migrationBuilder.DropColumn(
                name: "UpdatedAt",
                table: "Clinic");

            migrationBuilder.DropColumn(
                name: "UserCI",
                table: "Appointment");

            migrationBuilder.DropColumn(
                name: "UserCI",
                table: "AccountReceivable");
        }
    }
}
