namespace DataAccess;

public sealed class Transaction
{
    public ulong Id { get; set; }
    public WalletBalance TargetWallet { get; set; } = new WalletBalance();
    public WalletBalance SourceWallet { get; set; } = new WalletBalance();
    public TransactionStatus Status { get; set; }
    public TransactionType Type { get; set; }
    public decimal Amount { get; set; }
}