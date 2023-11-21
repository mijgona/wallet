namespace BusinessLogic;

public readonly record struct WalletInfo(
    long Id,
    float Balance,
    long UserId
);
