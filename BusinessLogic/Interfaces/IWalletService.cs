using DataAccess;

namespace BusinessLogic;

public interface IWalletService
{
    public Task<Wallet?> CreateWallet(WalletInfo userInfo);
}