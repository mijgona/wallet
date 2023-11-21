using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace DataAccess.Migrations
{
    /// <inheritdoc />
    public partial class RefMig : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateIndex(
                name: "IX_Transactions_SourceWalletId",
                table: "Transactions",
                column: "SourceWalletId");

            migrationBuilder.CreateIndex(
                name: "IX_Transactions_TargetWalletId",
                table: "Transactions",
                column: "TargetWalletId");

            migrationBuilder.AddForeignKey(
                name: "FK_Transactions_Wallet_SourceWalletId",
                table: "Transactions",
                column: "SourceWalletId",
                principalTable: "Wallet",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);

            migrationBuilder.AddForeignKey(
                name: "FK_Transactions_Wallet_TargetWalletId",
                table: "Transactions",
                column: "TargetWalletId",
                principalTable: "Wallet",
                principalColumn: "Id",
                onDelete: ReferentialAction.Cascade);
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropForeignKey(
                name: "FK_Transactions_Wallet_SourceWalletId",
                table: "Transactions");

            migrationBuilder.DropForeignKey(
                name: "FK_Transactions_Wallet_TargetWalletId",
                table: "Transactions");

            migrationBuilder.DropIndex(
                name: "IX_Transactions_SourceWalletId",
                table: "Transactions");

            migrationBuilder.DropIndex(
                name: "IX_Transactions_TargetWalletId",
                table: "Transactions");
        }
    }
}
