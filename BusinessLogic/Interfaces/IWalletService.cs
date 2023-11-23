using DataAccess;

namespace BusinessLogic;

public interface IWalletService
{
    public ValueTask<Wallet> CreateWalletAsync(WalletInfo userInfo, CancellationToken token);
}