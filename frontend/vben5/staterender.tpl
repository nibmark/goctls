
    {
      title: $t('common.status'),
      field: 'state',
      slots: {
        default: (e) =>
          h(Tag, {
            color: e.row.state ? 'success' : 'default',
          }, () => (e.row.state ? $t('common.on') : $t('common.off'))),
      },
    },
