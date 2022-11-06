from django.contrib import admin
from django.db.models import Q
from django.http import HttpRequest

from .models import DakuUser


# Register your models here.
@admin.register(DakuUser)
class BuyerAdmin(admin.ModelAdmin):

    def has_delete_permission(self, request: HttpRequest, obj = None) -> bool:
        if obj is None:
            return True
        else:
            return not obj.is_superuser

    def has_add_permission(self, request: HttpRequest) -> bool:
        return False

    def has_change_permission(self, request: HttpRequest, obj=None) -> bool:
        return False

    def get_queryset(self, request):
        return self.model.objects.filter(Q(is_active=True))

    ordering = ('-created',)
    list_display = (
        'user_name', 'email',
    )

    readonly_fields = (
        'user_name',
        'email',
    )

    fieldsets = (
        ('Credentials', {'fields': ('user_name', 'email',)}),
    )