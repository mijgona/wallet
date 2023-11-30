using DataAccess;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;

namespace BusinessLogic;

public static class ServiceExtension
{
    public static void ConfigureWalletServices(this IServiceCollection services, IConfiguration configuration)
    {
        services.ConfigureDataAcces(configuration);

        services.AddTransient<IUserService, UserService>();
        services.AddTransient<ITransactionService, TransactionService>();
        services.AddTransient<IWalletService, WalletService>();
        services.AddTransient<IWorkerService, WorkerService>();
    }
}