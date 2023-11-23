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

    public async ValueTask<Wallet> CreateWalletAsync(WalletInfo walletInfo, CancellationToken token)
    {
        Wallet newWallet = new()
        {
            UserId = walletInfo.UserId,
            Balance = walletInfo.Balance
        };
        return await _walletRepository.CreateAsync(newWallet, token) ?? throw new InvalidOperationException();
    }
}