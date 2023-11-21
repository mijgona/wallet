namespace DataAccess.Repositories;

public interface IWalletRepository
{
    ValueTask<Wallet?> CreateAsync(Wallet? wallet, CancellationToken token = default);
    ValueTask<Wallet?> GetWalletByUserId(long userId, CancellationToken token = default);
    
}