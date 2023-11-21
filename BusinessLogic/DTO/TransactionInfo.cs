namespace BusinessLogic;

public readonly record struct TransactionInfo(
    long Id,
    long SourceWalletId,
    long TargetWalletId,
    string Status,
    string Type,
    float Amount );