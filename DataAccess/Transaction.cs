namespace DataAccess;

public sealed class Transaction
{
    public long Id { get; set; }
    public long SourceWalletId { get; set; }
    public long TargetWalletId { get; set; }
    
    public Wallet? SourceWallet { get; set; }
    public Wallet? TargetWallet { get; set; }
    public TransactionStatus Status { get; set; } = TransactionStatus.Pending;
    public TransactionType Type { get; set; }
    public float Amount { get; set; }
}