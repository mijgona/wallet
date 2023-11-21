using DataAccess;

namespace BusinessLogic;

public interface IUserService
{
    public Task<User> CreateUser(UserInfo userInfo);
    public bool RemoveUser(string firstName, string lastName);
}
