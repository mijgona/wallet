using Microsoft.Extensions.Logging;

namespace BusinessLogic;

public class WorkerService : IWorkerService
{
    private readonly ITransactionService _transactionService;
    private readonly ILogger<IWorkerService> _logger;
    public WorkerService(ITransactionService transactionService, ILogger<IWorkerService> logger)
    {
        _transactionService = transactionService;
        _logger = logger;
    }

    public void StartWorker()
    {
        _logger.LogError("GOT ON BUSINESS LOGIC");
        _transactionService.FinalizeTransactionsAsync(new CancellationToken());
    }
}