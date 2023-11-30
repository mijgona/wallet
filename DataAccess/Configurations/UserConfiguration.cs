using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class UserConfiguration: IEntityTypeConfiguration<User>
{
    public void Configure(EntityTypeBuilder<User> modelBuilder)
    {
        modelBuilder
            .HasKey(p => p.Id)
            .HasName("pk_id");

        modelBuilder
            .HasIndex(p => p.PhoneNumber)
            .IsUnique();
    }
}