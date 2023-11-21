namespace DataAccess;

public class User
{
   public long Id { get; set; }
   public string Name { get; set;} = string.Empty;
   public string UserName { get; set;} = string.Empty;
   public string Password { get; set;} = string.Empty;
   public string LastName { get; set;} = string.Empty;
   public string? MiddleName { get; set;}
   public string PhoneNumber { get; set;} = string.Empty;
   public string? PassportNumber { get; set;} = string.Empty;
   public short Age { get; set; }
   public Gender Gender { get; set; }
   public Wallet? Balance { get; set; }
}