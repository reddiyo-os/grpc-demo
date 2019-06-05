all:

    @echo "Text from env. var.: $(FOO)"
    @echo ""

    @while [ -z "$$CONTINUE" ]; do \
        read -r -p "Type anything but Y or y to exit. [y/N]: " CONTINUE; \
    done ; \
    [ $$CONTINUE = "y" ] || [ $$CONTINUE = "Y" ] || (echo "Exiting."; exit 1;)
    @echo "..do more.."