using DataAccess;

namespace BusinessLogic;

public interface IUserService
{
    public Task<User> CreateUser(UserInfo userInfo);
    public ValueTask<User> GetUserByUsername(string userName);
}
