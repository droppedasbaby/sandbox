"""
Implementation of AutoComplete.
"""


class AutoComplete:
    """
    Implementation of AutoComplete.
    """

    def __init__(self, data_structure_type, len_limit=100):
        """
        Create an instance of the data structure that will be used as the storage for this class.

        :param data_structure_type:
        :return:
        """
        self.data_store = data_structure_type()
        self.len_limit = len_limit

    def insert(self, phrases):
        """
        Add all the phrases to the data store.

        :param phrases:
        :return:
        """
        if isinstance(phrases, list):
            for phrase in phrases:
                if len(phrase) > self.len_limit:
                    raise ValueError(
                        "AutoComplete::add len(phrase) > len_limit defined, phrase: "
                        + phrase
                    )
                else:
                    self.data_store.insert(phrase.lower().rstrip())
        else:
            raise TypeError(
                "AutoComplete::add phrases is of incorrect type, it should be of type list."
            )

    def query(self, prefix):
        """
        Query the data store for the phrases matching the prefix and return the results as an array.

        :param prefix:
        :return:
        """
        return self.data_store.query(prefix.lower())
