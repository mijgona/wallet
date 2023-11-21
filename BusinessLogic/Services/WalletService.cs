using DataAccess;
using DataAccess.Repositories;

namespace BusinessLogic;

public class WalletService : IWalletService
{
    private readonly IWalletRepository _walletRepository;

    public WalletService(IWalletRepository walletRepository)
    {
        _walletRepository = walletRepository;
    }

    public async Task<Wallet?> CreateWallet(WalletInfo walletInfo)
    {
        Wallet newWallet = new()
        {
            UserId = walletInfo.UserId,
            Balance = walletInfo.Balance
        };
        return await _walletRepository.CreateAsync(newWallet);
    }
}