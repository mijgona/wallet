namespace DataAccess.Repositories;

public interface IWalletRepository
{
    ValueTask<Wallet?> CreateAsync(Wallet? wallet, CancellationToken token = default);
    
}