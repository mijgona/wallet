using Microsoft.EntityFrameworkCore;

namespace DataAccess;

public class WalletBalance
{
    public ulong Id { get; set;}
    public decimal Balance { get; set; }
    public ulong UserId { get; set; }
    public User User { get; set; }
    public List<Transaction> Transactions { get; set; }
}