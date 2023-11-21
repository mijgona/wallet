using DataAccess;

namespace BusinessLogic;

public sealed class TransactionService : ITransactionService
{
    private readonly ITransactionRepository _transactionRepository;

    public TransactionService(ITransactionRepository transactionRepository)
    {
        _transactionRepository = transactionRepository;
    }

    public ValueTask<Transaction> CreateTransactionAsync(TransactionInfo transaction, CancellationToken token)
    {
        
        Transaction newTransaction = new()
        {
            SourceWalletId = transaction.SourceWalletId,
            TargetWalletId = transaction.TargetWalletId,
            Status = TransactionStatus.Pending,
            Type = transaction.Type.ToTransactionTypeEnum(),
            Amount = transaction.Amount,
        };
        
        var result =_transactionRepository.CreateAsync(newTransaction, token);
        
        return result;
    }
    
    
    public async ValueTask<List<Transaction>> GetTransactionsAsync(long userId, CancellationToken token)
    {

        return await _transactionRepository.GetTransactionsAsync(userId, default);;
    }
}
