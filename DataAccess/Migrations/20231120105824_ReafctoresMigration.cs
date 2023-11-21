using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace DataAccess.Migrations
{
    /// <inheritdoc />
    public partial class ReafctoresMigration : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Transactions_Wallet_WalletBalanceId",
                table: "Transactions");

            migrationBuilder.DropIndex(
                name: "IX_Transactions_WalletBalanceId",
                table: "Transactions");

            migrationBuilder.AddColumn<long>(
                name: "WalletId",
                table: "Transactions",
                type: "bigint",
                nullable: false,
                defaultValue: 0L);

            migrationBuilder.CreateIndex(
                name: "IX_Transactions_WalletId",
                table: "Transactions",
                column: "WalletId");

            migrationBuilder.AddForeignKey(
                name: "FK_Transactions_Wallet_WalletId",
                table: "Transactions",
                column: "WalletId",
                principalTable: "Wallet",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Transactions_Wallet_WalletId",
                table: "Transactions");

            migrationBuilder.DropIndex(
                name: "IX_Transactions_WalletId",
                table: "Transactions");

            migrationBuilder.DropColumn(
                name: "WalletId",
                table: "Transactions");

            migrationBuilder.CreateIndex(
                name: "IX_Transactions_WalletBalanceId",
                table: "Transactions",
                column: "WalletBalanceId");

            migrationBuilder.AddForeignKey(
                name: "FK_Transactions_Wallet_WalletBalanceId",
                table: "Transactions",
                column: "WalletBalanceId",
                principalTable: "Wallet",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }
    }
}
