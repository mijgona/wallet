using DataAccess;

namespace BusinessLogic;

public interface IWalletService
{
    public ValueTask<Wallet> CreateWalletAsync(WalletInfo userInfo, CancellationToken token);
    public ValueTask<Wallet> GetWalletByUserIdAsync(long userId, CancellationToken token);
}