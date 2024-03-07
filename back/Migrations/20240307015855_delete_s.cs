using System;
using Microsoft.EntityFrameworkCore.Migrations;
using Npgsql.EntityFrameworkCore.PostgreSQL.Metadata;

#nullable disable

namespace back.Migrations
{
    /// <inheritdoc />
    public partial class delete_s : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "AccountReceivableDetails");

            migrationBuilder.AddColumn<DateTime>(
                name: "CreatedAt",
                table: "Price",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.AddColumn<DateTime>(
                name: "UpdatedAt",
                table: "Price",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.AddColumn<string>(
                name: "Code",
                table: "AccountReceivable",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<DateTime>(
                name: "CreatedAt",
                table: "AccountReceivable",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.AddColumn<string>(
                name: "Status",
                table: "AccountReceivable",
                type: "text",
                nullable: false,
                defaultValue: "");

            migrationBuilder.AddColumn<DateTime>(
                name: "UpdatedAt",
                table: "AccountReceivable",
                type: "timestamp with time zone",
                nullable: false,
                defaultValue: new DateTime(1, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified));

            migrationBuilder.CreateTable(
                name: "AccountReceivableDetail",
                columns: table => new
                {
                    ID = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    AccountReceivableID = table.Column<int>(type: "integer", nullable: false),
                    Amount = table.Column<double>(type: "double precision", nullable: false),
                    Status = table.Column<string>(type: "text", nullable: false),
                    CreatedAt = table.Column<DateTime>(type: "timestamp with time zone", nullable: false),
                    UpdatedAt = table.Column<DateTime>(type: "timestamp with time zone", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_AccountReceivableDetail", x => x.ID);
                    table.ForeignKey(
                        name: "FK_AccountReceivableDetail_AccountReceivable_AccountReceivable~",
                        column: x => x.AccountReceivableID,
                        principalTable: "AccountReceivable",
                        principalColumn: "ID",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_AccountReceivableDetail_AccountReceivableID",
                table: "AccountReceivableDetail",
                column: "AccountReceivableID");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "AccountReceivableDetail");

            migrationBuilder.DropColumn(
                name: "CreatedAt",
                table: "Price");

            migrationBuilder.DropColumn(
                name: "UpdatedAt",
                table: "Price");

            migrationBuilder.DropColumn(
                name: "Code",
                table: "AccountReceivable");

            migrationBuilder.DropColumn(
                name: "CreatedAt",
                table: "AccountReceivable");

            migrationBuilder.DropColumn(
                name: "Status",
                table: "AccountReceivable");

            migrationBuilder.DropColumn(
                name: "UpdatedAt",
                table: "AccountReceivable");

            migrationBuilder.CreateTable(
                name: "AccountReceivableDetails",
                columns: table => new
                {
                    ID = table.Column<int>(type: "integer", nullable: false)
                        .Annotation("Npgsql:ValueGenerationStrategy", NpgsqlValueGenerationStrategy.IdentityByDefaultColumn),
                    AccountReceivableID = table.Column<int>(type: "integer", nullable: false),
                    Amount = table.Column<double>(type: "double precision", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_AccountReceivableDetails", x => x.ID);
                    table.ForeignKey(
                        name: "FK_AccountReceivableDetails_AccountReceivable_AccountReceivabl~",
                        column: x => x.AccountReceivableID,
                        principalTable: "AccountReceivable",
                        principalColumn: "ID",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_AccountReceivableDetails_AccountReceivableID",
                table: "AccountReceivableDetails",
                column: "AccountReceivableID");
        }
    }
}
