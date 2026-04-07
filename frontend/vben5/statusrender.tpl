
    {
      title: $t('common.status'),
      field: 'status',
      slots: {
        default: (e) =>
          h(Tag, {
            color: e.row.status === 1 ? 'success' : 'default',
          }, () =>
            e.row.status === 1 ? $t('common.on') : $t('common.off')),
      },
    },
