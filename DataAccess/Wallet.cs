namespace DataAccess;

public class Wallet
{
    public long Id { get; set;}
    public float Balance { get; set; }
    public long UserId { get; set; }
    public User? User { get; set; }
}